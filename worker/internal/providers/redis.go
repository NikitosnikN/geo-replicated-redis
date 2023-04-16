package providers

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"strconv"
	"time"
	"worker/internal/common"
)

type RedisProvider struct {
	Client *redis.Client
}

func (r *RedisProvider) ProceedCommand(cmd common.CommandV1) {
	var ctx = context.Background()

	switch cmd.Command {
	case "SET":
		fmt.Printf("SET %s\n", cmd.Args)

		if len(cmd.Args) < 3 {
			fmt.Printf("invalid args for SET command %s\n", cmd.Args)
			return
		}

		duration, _ := strconv.Atoi(cmd.Args[2])
		r.Client.Set(ctx, cmd.Args[0], cmd.Args[1], time.Duration(duration)*time.Second)
	case "DEL":
		fmt.Printf("DEL %s\n", cmd.Args)
		r.Client.Del(ctx, cmd.Args...)
	default:
		fmt.Printf("Command %s not supported\n", cmd.Command)
	}
}
