package mappers

import (
	data "github.com/cassiusbessa/vision-social-media/data-access/sqlc-config"
	"github.com/cassiusbessa/vision-social-media/domain/core/entities"
)

func ReactionEntityToCreateReactionQueryParams(reaction *entities.Reaction) data.CreateReactionParams {
	return data.CreateReactionParams{
		ID:           reaction.ID,
		PostID:       reaction.PostID,
		CommentID:    reaction.ParentID,
		UserID:       reaction.UserID,
		ReactionType: string(reaction.Type),
		CreatedAt:    reaction.CreatedAt,
	}
}

func ReactionDbEntityToReaction(reaction data.Reaction) *entities.Reaction {

	var reactionType entities.ReactionType

	switch reaction.ReactionType {
	case "Like":
		reactionType = entities.Like
	case "Dislike":
		reactionType = entities.Dislike
	case "Love":
		reactionType = entities.Love
	case "Wow":
		reactionType = entities.Wow
	case "Angry":
		reactionType = entities.Angry
	}

	return entities.NewReaction(
		entities.ReactionWithID(reaction.ID),
		entities.ReactionWithPostID(reaction.PostID),
		entities.ReactionWithUserID(reaction.UserID),
		entities.ReactionWithReactionType(reactionType),
	)
}
