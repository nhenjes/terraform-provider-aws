// Code generated by "internal/generate/listpages/main.go -AWSSDKVersion=2 -ListOps=ListActivatedRulesInRuleGroup,ListByteMatchSets,ListGeoMatchSets,ListIPSets,ListRateBasedRules,ListRegexMatchSets,ListRegexPatternSets,ListRuleGroups,ListRules,ListSizeConstraintSets,ListSqlInjectionMatchSets,ListWebACLs,ListXssMatchSets -Paginator=NextMarker"; DO NOT EDIT.

package waf

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/waf"
)

func listActivatedRulesInRuleGroupPages(ctx context.Context, conn *waf.Client, input *waf.ListActivatedRulesInRuleGroupInput, fn func(*waf.ListActivatedRulesInRuleGroupOutput, bool) bool) error {
	for {
		output, err := conn.ListActivatedRulesInRuleGroup(ctx, input)
		if err != nil {
			return err
		}

		lastPage := aws.ToString(output.NextMarker) == ""
		if !fn(output, lastPage) || lastPage {
			break
		}

		input.NextMarker = output.NextMarker
	}
	return nil
}
func listByteMatchSetsPages(ctx context.Context, conn *waf.Client, input *waf.ListByteMatchSetsInput, fn func(*waf.ListByteMatchSetsOutput, bool) bool) error {
	for {
		output, err := conn.ListByteMatchSets(ctx, input)
		if err != nil {
			return err
		}

		lastPage := aws.ToString(output.NextMarker) == ""
		if !fn(output, lastPage) || lastPage {
			break
		}

		input.NextMarker = output.NextMarker
	}
	return nil
}
func listGeoMatchSetsPages(ctx context.Context, conn *waf.Client, input *waf.ListGeoMatchSetsInput, fn func(*waf.ListGeoMatchSetsOutput, bool) bool) error {
	for {
		output, err := conn.ListGeoMatchSets(ctx, input)
		if err != nil {
			return err
		}

		lastPage := aws.ToString(output.NextMarker) == ""
		if !fn(output, lastPage) || lastPage {
			break
		}

		input.NextMarker = output.NextMarker
	}
	return nil
}
func listIPSetsPages(ctx context.Context, conn *waf.Client, input *waf.ListIPSetsInput, fn func(*waf.ListIPSetsOutput, bool) bool) error {
	for {
		output, err := conn.ListIPSets(ctx, input)
		if err != nil {
			return err
		}

		lastPage := aws.ToString(output.NextMarker) == ""
		if !fn(output, lastPage) || lastPage {
			break
		}

		input.NextMarker = output.NextMarker
	}
	return nil
}
func listRateBasedRulesPages(ctx context.Context, conn *waf.Client, input *waf.ListRateBasedRulesInput, fn func(*waf.ListRateBasedRulesOutput, bool) bool) error {
	for {
		output, err := conn.ListRateBasedRules(ctx, input)
		if err != nil {
			return err
		}

		lastPage := aws.ToString(output.NextMarker) == ""
		if !fn(output, lastPage) || lastPage {
			break
		}

		input.NextMarker = output.NextMarker
	}
	return nil
}
func listRegexMatchSetsPages(ctx context.Context, conn *waf.Client, input *waf.ListRegexMatchSetsInput, fn func(*waf.ListRegexMatchSetsOutput, bool) bool) error {
	for {
		output, err := conn.ListRegexMatchSets(ctx, input)
		if err != nil {
			return err
		}

		lastPage := aws.ToString(output.NextMarker) == ""
		if !fn(output, lastPage) || lastPage {
			break
		}

		input.NextMarker = output.NextMarker
	}
	return nil
}
func listRegexPatternSetsPages(ctx context.Context, conn *waf.Client, input *waf.ListRegexPatternSetsInput, fn func(*waf.ListRegexPatternSetsOutput, bool) bool) error {
	for {
		output, err := conn.ListRegexPatternSets(ctx, input)
		if err != nil {
			return err
		}

		lastPage := aws.ToString(output.NextMarker) == ""
		if !fn(output, lastPage) || lastPage {
			break
		}

		input.NextMarker = output.NextMarker
	}
	return nil
}
func listRuleGroupsPages(ctx context.Context, conn *waf.Client, input *waf.ListRuleGroupsInput, fn func(*waf.ListRuleGroupsOutput, bool) bool) error {
	for {
		output, err := conn.ListRuleGroups(ctx, input)
		if err != nil {
			return err
		}

		lastPage := aws.ToString(output.NextMarker) == ""
		if !fn(output, lastPage) || lastPage {
			break
		}

		input.NextMarker = output.NextMarker
	}
	return nil
}
func listRulesPages(ctx context.Context, conn *waf.Client, input *waf.ListRulesInput, fn func(*waf.ListRulesOutput, bool) bool) error {
	for {
		output, err := conn.ListRules(ctx, input)
		if err != nil {
			return err
		}

		lastPage := aws.ToString(output.NextMarker) == ""
		if !fn(output, lastPage) || lastPage {
			break
		}

		input.NextMarker = output.NextMarker
	}
	return nil
}
func listSizeConstraintSetsPages(ctx context.Context, conn *waf.Client, input *waf.ListSizeConstraintSetsInput, fn func(*waf.ListSizeConstraintSetsOutput, bool) bool) error {
	for {
		output, err := conn.ListSizeConstraintSets(ctx, input)
		if err != nil {
			return err
		}

		lastPage := aws.ToString(output.NextMarker) == ""
		if !fn(output, lastPage) || lastPage {
			break
		}

		input.NextMarker = output.NextMarker
	}
	return nil
}
func listSQLInjectionMatchSetsPages(ctx context.Context, conn *waf.Client, input *waf.ListSqlInjectionMatchSetsInput, fn func(*waf.ListSqlInjectionMatchSetsOutput, bool) bool) error {
	for {
		output, err := conn.ListSqlInjectionMatchSets(ctx, input)
		if err != nil {
			return err
		}

		lastPage := aws.ToString(output.NextMarker) == ""
		if !fn(output, lastPage) || lastPage {
			break
		}

		input.NextMarker = output.NextMarker
	}
	return nil
}
func listWebACLsPages(ctx context.Context, conn *waf.Client, input *waf.ListWebACLsInput, fn func(*waf.ListWebACLsOutput, bool) bool) error {
	for {
		output, err := conn.ListWebACLs(ctx, input)
		if err != nil {
			return err
		}

		lastPage := aws.ToString(output.NextMarker) == ""
		if !fn(output, lastPage) || lastPage {
			break
		}

		input.NextMarker = output.NextMarker
	}
	return nil
}
func listXSSMatchSetsPages(ctx context.Context, conn *waf.Client, input *waf.ListXssMatchSetsInput, fn func(*waf.ListXssMatchSetsOutput, bool) bool) error {
	for {
		output, err := conn.ListXssMatchSets(ctx, input)
		if err != nil {
			return err
		}

		lastPage := aws.ToString(output.NextMarker) == ""
		if !fn(output, lastPage) || lastPage {
			break
		}

		input.NextMarker = output.NextMarker
	}
	return nil
}
