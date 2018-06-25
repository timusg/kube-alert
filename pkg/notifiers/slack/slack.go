package slack

import (
	"github.com/bluele/slack"
	"github.com/bpineau/kube-alert/config"
)

// Notifier implements notifers.Notifier for Datadog
type Notifier struct {
}

var tags = []string{
	"context:kubernetes",
	"origin:kube-alert",
}

// Notify sends alerts as Datadog events
func (l *Notifier) Notify(c *config.AlertConfig, title string, msg string) error {
	if c.SlackToken == "" {
		c.Logger.Debug("Omitting Slack notification, api key missing")
		return nil
	}

	api := slack.New(c.SlackToken)

	err := api.ChatPostMessage(c.SlackChannelName, title+" -> "+msg, nil)

	if err != nil {
		c.Logger.Warning("failed to post to slack: %s", err)
	}

	return err
}
