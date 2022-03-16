package learning_content

import (
	"context"

	learning_content_client "github.com/1024pix/go-learning-content-client"
	"github.com/turbot/steampipe-plugin-sdk/v2/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v2/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v2/plugin/transform"
)

func skillsTable() *plugin.Table {
	return &plugin.Table{
		Name: "skills",
		List: &plugin.ListConfig{
			Hydrate: hydrateSkillsList,
		},
		DefaultTransform: transform.FromJSONTag(),
		Columns: []*plugin.Column{
			{Name: "id", Type: proto.ColumnType_STRING},
			{Name: "name", Type: proto.ColumnType_STRING},
			{Name: "hintFrFr", Type: proto.ColumnType_STRING},
			{Name: "hintEnUs", Type: proto.ColumnType_STRING},
			{Name: "hintStatus", Type: proto.ColumnType_STRING},
			// FIXME tutorialIds
			// FIXME learningMoreTutorialIds
			{Name: "pixValue", Type: proto.ColumnType_INT},
			{Name: "competenceId", Type: proto.ColumnType_STRING},
			{Name: "status", Type: proto.ColumnType_STRING},
			{Name: "tubeId", Type: proto.ColumnType_STRING},
			{Name: "version", Type: proto.ColumnType_STRING},
		},
	}
}

func hydrateSkillsList(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	logger.Debug("hydrating skills list")

	config := getConfig(d.Connection)

	c := learning_content_client.New(
		learning_content_client.WithApiURL(*config.ApiURL),
		learning_content_client.WithApiKey(*config.ApiKey),
	)

	r, err := c.GetLatestRelease()
	if err != nil {
		return nil, err
	}

	for _, skill := range r.Content.Skills {
		d.StreamListItem(ctx, skill)
	}

	return nil, nil
}
