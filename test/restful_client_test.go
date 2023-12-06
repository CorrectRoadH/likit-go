package test

import (
	"context"
	"testing"

	"github.com/CorrectRoadH/likit-go"
	"github.com/stretchr/testify/assert"
)

func TestVote(t *testing.T) {
	// create vote server
	voteServer := likit.NewVoteServer("https://likit.zeabur.app")

	// vote
	count, err := voteServer.Vote(context.Background(), "COMMENT_LIKE", "123", "userIdqw")
	assert.NoError(t, err)

	// check count
	assert.Equal(t, int64(1), count)
}
