package record

import (
	"errors"
	"io"
	"sort"
	"strconv"
	"sync"
	"time"
)

const (
	NONE = "NONE"
	FIFO = "FIFO"
	HARD = "HARD"
)

type Record struct {
	Timestamp   int64             // UNIX timestamp in microseconds
	Size        int               // Size of data
	Last        bool              // Last record in the query
	ContentType string            // Content type of data
	Labels      map[string]string // Labels of record
	Data        []byte            // Data content
}

func (r *Record) GetDatetime() time.Time {
	return time.Unix(0, r.Timestamp*1000) // Convert microseconds to nanoseconds
}

type Batch struct {
	records    map[int64]*Record
	totalSize  int
	lastAccess time.Time
	mu         sync.Mutex
}

func NewBatch() *Batch {
	return &Batch{
		records:    make(map[int64]*Record),
		totalSize:  0,
		lastAccess: time.Now(),
	}
}

func (b *Batch) Add(timestamp interface{}, data []byte, contentType string, labels map[string]string) error {
	b.mu.Lock()
	defer b.mu.Unlock()

	if labels == nil {
		labels = make(map[string]string)
	}

	var ts int64
	switch v := timestamp.(type) {
	case int64:
		ts = v
	case time.Time:
		ts = v.UnixNano() / 1000 // Convert nanoseconds to microseconds
	case float64:
		ts = int64(v * 1e6) // Convert seconds to microseconds
	case string:
		t, err := time.Parse(time.RFC3339, v)
		if err != nil {
			return err
		}
		ts = t.UnixNano() / 1000
	default:
		return errors.New("unsupported timestamp type")
	}

	record := &Record{
		Timestamp:   ts,
		Size:        len(data),
		ContentType: contentType,
		Labels:      labels,
		Data:        data,
	}

	b.records[ts] = record
	b.totalSize += len(data)
	b.lastAccess = time.Now()

	return nil
}

func (b *Batch) Items() []*Record {
	b.mu.Lock()
	defer b.mu.Unlock()

	var items []*Record
	for _, record := range b.records {
		items = append(items, record)
	}

	sort.Slice(items, func(i, j int) bool {
		return items[i].Timestamp < items[j].Timestamp
	})

	return items
}

func (b *Batch) Size() int {
	b.mu.Lock()
	defer b.mu.Unlock()
	return b.totalSize
}

func (b *Batch) LastAccess() time.Time {
	b.mu.Lock()
	defer b.mu.Unlock()
	return b.lastAccess
}

func (b *Batch) Clear() {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.records = make(map[int64]*Record)
	b.totalSize = 0
	b.lastAccess = time.Now()
}

func (b *Batch) Len() int {
	b.mu.Lock()
	defer b.mu.Unlock()
	return len(b.records)
}

func ParseRecord(headers map[string]string, body io.Reader, last bool) (*Record, error) {
	timestamp, err := parseHeaderInt64(headers, "x-reduct-time")
	if err != nil {
		return nil, err
	}

	size, err := parseHeaderInt64(headers, "content-length")
	if err != nil {
		return nil, err
	}

	contentType := headers["content-type"]
	if contentType == "" {
		contentType = "application/octet-stream"
	}

	labels := make(map[string]string)
	for key, value := range headers {
		if len(key) > len(LABEL_PREFIX) && key[:len(LABEL_PREFIX)] == LABEL_PREFIX {
			labels[key[len(LABEL_PREFIX):]] = value
		}
	}

	data, err := io.ReadAll(body)
	if err != nil {
		return nil, err
	}

	return &Record{
		Timestamp:   timestamp,
		Size:        int(size),
		Last:        last,
		ContentType: contentType,
		Labels:      labels,
		Data:        data,
	}, nil
}

func parseHeaderInt64(headers map[string]string, key string) (int64, error) {
	value, ok := headers[key]
	if !ok {
		return 0, errors.New("missing header: " + key)
	}
	return strconv.ParseInt(value, 10, 64)
}

const (
	LABEL_PREFIX = "x-reduct-label-"
	TIME_PREFIX  = "x-reduct-time-"
	ERROR_PREFIX = "x-reduct-error-"
	CHUNK_SIZE   = 16000
)
