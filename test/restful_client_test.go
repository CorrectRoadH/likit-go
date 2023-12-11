package test

import (
	"context"
	"strings"
	"testing"

	"math/rand"
	"time"

	"github.com/CorrectRoadH/likit-go"
	"github.com/stretchr/testify/assert"
)

func randomName() string {
	rand.Seed(time.Now().UnixNano())
	chars := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	length := 10
	var b strings.Builder
	for i := 0; i < length; i++ {
		b.WriteRune(chars[rand.Intn(len(chars))])
	}
	return b.String()
}
func TestVote(t *testing.T) {
	// create vote server
	voteServer := likit.NewVoteServer("localhost:4778")

	// random messageId
	messageId := randomName()
	userId := randomName()

	// vote
	count, err := voteServer.Vote(context.Background(), "COMMENT_LIKE", messageId, userId)
	assert.NoError(t, err)

	// check count
	assert.Equal(t, int64(1), count)
}

func TestIsVotd(t *testing.T) {
	// create vote server
	voteServer := likit.NewVoteServer("localhost:4778")

	// random messageId
	messageId := randomName()
	userId := randomName()
	// vote
	isVoted, err := voteServer.IsVote(context.Background(), "COMMENT_LIKE", messageId, userId)
	assert.NoError(t, err)
	assert.Equal(t, false, isVoted)

	// vote
	count, err := voteServer.Vote(context.Background(), "COMMENT_LIKE", messageId, userId)
	assert.NoError(t, err)
	assert.Equal(t, int64(1), count)

	isVoted, err = voteServer.IsVote(context.Background(), "COMMENT_LIKE", messageId, userId)
	assert.NoError(t, err)
	assert.Equal(t, true, isVoted)
}
