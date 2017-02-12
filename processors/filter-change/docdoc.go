// Code generated by "bitfanDoc "; DO NOT EDIT
package change

import "github.com/vjeantet/bitfan/processors/doc"

func (p *processor) Doc() *doc.Processor {
	return &doc.Processor{
  Name:     "change",
  Doc:      "This rule will monitor a certain field and match if that field changes. The field must change with respect to the last event",
  DocShort: "drop event when field value is the same in the last event",
  Options:  &doc.ProcessorOptions{
    Doc:     "",
    Options: []*doc.ProcessorOption{
      &doc.ProcessorOption{
        Name:         "AddField",
        Alias:        "add_field",
        Doc:          "If this filter is successful, add any arbitrary fields to this event.",
        Required:     false,
        Type:         "hash",
        DefaultValue: nil,
        ExampleLS:    "",
      },
      &doc.ProcessorOption{
        Name:         "AddTag",
        Alias:        "add_tag",
        Doc:          "If this filter is successful, add arbitrary tags to the event. Tags can be dynamic\nand include parts of the event using the %{field} syntax.",
        Required:     false,
        Type:         "array",
        DefaultValue: nil,
        ExampleLS:    "",
      },
      &doc.ProcessorOption{
        Name:         "RemoveField",
        Alias:        "remove_field",
        Doc:          "If this filter is successful, remove arbitrary fields from this event. Example:\n` kv {\n`   remove_field => [ \"foo_%{somefield}\" ]\n` }",
        Required:     false,
        Type:         "array",
        DefaultValue: nil,
        ExampleLS:    "",
      },
      &doc.ProcessorOption{
        Name:         "RemoveTag",
        Alias:        "remove_tag",
        Doc:          "If this filter is successful, remove arbitrary tags from the event. Tags can be dynamic and include parts of the event using the %{field} syntax.\nExample:\n` kv {\n`   remove_tag => [ \"foo_%{somefield}\" ]\n` }\nIf the event has field \"somefield\" == \"hello\" this filter, on success, would remove the tag foo_hello if it is present. The second example would remove a sad, unwanted tag as well.",
        Required:     false,
        Type:         "array",
        DefaultValue: nil,
        ExampleLS:    "",
      },
      &doc.ProcessorOption{
        Name:         "CompareField",
        Alias:        "compare_field",
        Doc:          "The name of the field to use to compare to the blacklist.\nIf the field is null, those events will be ignored.",
        Required:     true,
        Type:         "string",
        DefaultValue: nil,
        ExampleLS:    "compare_field => \"message\"",
      },
      &doc.ProcessorOption{
        Name:         "IgnoreNull",
        Alias:        "ignore_null",
        Doc:          "If true, events without a compare_key field will not count as changed.",
        Required:     false,
        Type:         "bool",
        DefaultValue: "true",
        ExampleLS:    "",
      },
    },
  },
  Ports: []*doc.ProcessorPort{
    &doc.ProcessorPort{
      Default: true,
      Name:    "PORT_SUCCESS",
      Number:  0,
      Doc:     "",
    },
  },
}
}