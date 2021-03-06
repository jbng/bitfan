// Code generated by "bitfanDoc "; DO NOT EDIT
package httpoutprocessor

import "github.com/vjeantet/bitfan/processors/doc"

func (p *processor) Doc() *doc.Processor {
	return &doc.Processor{
  Name:       "httpoutprocessor",
  ImportPath: "github.com/vjeantet/bitfan/processors/httpout",
  Doc:        "Display on http the last received event\n\nURL is available as http://webhookhost/pipelineName/pluginLabel/URI\n\n* webhookhost is defined by bitfan at startup\n* pluginLabel is defined in pipeline configuration, it's the named processor if you put one, or `httpout` by default\n* URI is defined in plugin configuration (see below)",
  DocShort:   "Reads events from standard input",
  Options:    &doc.ProcessorOptions{
    Doc:     "",
    Options: []*doc.ProcessorOption{
      &doc.ProcessorOption{
        Name:           "Codec",
        Alias:          "",
        Doc:            "The codec used for input data. Input codecs are a convenient method for decoding\nyour data before it enters the input, without needing a separate filter in your bitfan pipeline",
        Required:       false,
        Type:           "codec",
        DefaultValue:   "\"json\"",
        PossibleValues: []string{},
        ExampleLS:      "",
      },
      &doc.ProcessorOption{
        Name:           "Uri",
        Alias:          "",
        Doc:            "URI path",
        Required:       false,
        Type:           "string",
        DefaultValue:   "\"out\"",
        PossibleValues: []string{},
        ExampleLS:      "",
      },
      &doc.ProcessorOption{
        Name:           "Headers",
        Alias:          "",
        Doc:            "Add headers to output",
        Required:       false,
        Type:           "hash",
        DefaultValue:   nil,
        PossibleValues: []string{},
        ExampleLS:      "",
      },
    },
  },
  Ports: []*doc.ProcessorPort{},
}
}