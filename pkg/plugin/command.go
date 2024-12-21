package plugin

import "github.com/google/shlex"

type CommandContext struct {
	*MessageContext
	Tokens []string
}

func buildCommand(ctx *MessageContext) (*CommandContext, error) {
	tokens, err := shlex.Split(ctx.String()[len(bot.Config.CommandPrefix):])
	if err != nil {
		return nil, err
	}
	return &CommandContext{
		MessageContext: ctx,
		Tokens:         tokens,
	}, nil
}

func buildCommandHandle(prefix string) *EventHandle[*CommandContext] {
	tokens, err := shlex.Split(prefix)
	if err != nil {
		bot.Logger.Errorln("Failed to parse command prefix: ", err)
		return nil
	}

	return &EventHandle[*CommandContext]{
		matcher: func(ctx *CommandContext) bool {
			for i, token := range tokens {
				if len(ctx.Tokens) <= i || ctx.Tokens[i] != token {
					return false
				}
			}
			return true
		},
	}
}
