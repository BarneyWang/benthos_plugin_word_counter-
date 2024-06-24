package processor

import (
	"context"
	"encoding/json"

	"github.com/benthosdev/benthos/v4/public/service"
)

func init() {
	configSpec := service.NewConfigSpec()

	constructor := func(conf *service.ParsedConfig, mgr *service.Resources) (service.Processor, error) {

		return newWordCounterProcessor(mgr.Logger()), nil
		// return newWordCounterProcessor(mgr.Logger(), mgr.Metrics()), nil
	}
	err := service.RegisterProcessor("wordcounter", configSpec, constructor)
	if err != nil {
		panic(err)
	}

}

//------------------------------------------------------------------------------

type wordCounterProcessor struct {
	logger *service.Logger
}

type newMessage struct {
	Word  string `json:"word" yaml:"word"`
	Count int    `json:"count" yaml:"count"`
}

func newWordCounterProcessor(logger *service.Logger) *wordCounterProcessor {
	return &wordCounterProcessor{
		logger: logger,
	}
}

func (w *wordCounterProcessor) Process(ctx context.Context, m *service.Message) (service.MessageBatch, error) {
	bytesContent, err := m.AsBytes()
	if err != nil {
		return nil, err
	}
	word := string(bytesContent)
	count := len(word)

	w.logger.Infof("Woah! This length is like totally : %d", count)
	nm := newMessage{
		Word:  word,
		Count: count,
	}

	b, err := json.Marshal(nm)
	if err != nil {
		return nil, err
	}
	m.SetBytes(b)
	return []*service.Message{m}, nil
}

func (r *wordCounterProcessor) Close(ctx context.Context) error {
	return nil
}
