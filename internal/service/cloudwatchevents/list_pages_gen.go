// Code generated by "internal/generate/listpages/main.go -ListOps=ListEventBuses,ListRules,ListTargetsByRule"; DO NOT EDIT.

package cloudwatchevents

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/eventbridge"
)

func listEventBusesPages(conn *eventbridge.EventBridge, input *eventbridge.ListEventBusesInput, fn func(*eventbridge.ListEventBusesOutput, bool) bool) error {
	return listEventBusesPagesWithContext(context.Background(), conn, input, fn)
}

func listEventBusesPagesWithContext(ctx context.Context, conn *eventbridge.EventBridge, input *eventbridge.ListEventBusesInput, fn func(*eventbridge.ListEventBusesOutput, bool) bool) error {
	for {
		output, err := conn.ListEventBusesWithContext(ctx, input)
		if err != nil {
			return err
		}

		lastPage := aws.StringValue(output.NextToken) == ""
		if !fn(output, lastPage) || lastPage {
			break
		}

		input.NextToken = output.NextToken
	}
	return nil
}

func listRulesPages(conn *eventbridge.EventBridge, input *eventbridge.ListRulesInput, fn func(*eventbridge.ListRulesOutput, bool) bool) error {
	return listRulesPagesWithContext(context.Background(), conn, input, fn)
}

func listRulesPagesWithContext(ctx context.Context, conn *eventbridge.EventBridge, input *eventbridge.ListRulesInput, fn func(*eventbridge.ListRulesOutput, bool) bool) error {
	for {
		output, err := conn.ListRulesWithContext(ctx, input)
		if err != nil {
			return err
		}

		lastPage := aws.StringValue(output.NextToken) == ""
		if !fn(output, lastPage) || lastPage {
			break
		}

		input.NextToken = output.NextToken
	}
	return nil
}

func listTargetsByRulePages(conn *eventbridge.EventBridge, input *eventbridge.ListTargetsByRuleInput, fn func(*eventbridge.ListTargetsByRuleOutput, bool) bool) error {
	return listTargetsByRulePagesWithContext(context.Background(), conn, input, fn)
}

func listTargetsByRulePagesWithContext(ctx context.Context, conn *eventbridge.EventBridge, input *eventbridge.ListTargetsByRuleInput, fn func(*eventbridge.ListTargetsByRuleOutput, bool) bool) error {
	for {
		output, err := conn.ListTargetsByRuleWithContext(ctx, input)
		if err != nil {
			return err
		}

		lastPage := aws.StringValue(output.NextToken) == ""
		if !fn(output, lastPage) || lastPage {
			break
		}

		input.NextToken = output.NextToken
	}
	return nil
}
