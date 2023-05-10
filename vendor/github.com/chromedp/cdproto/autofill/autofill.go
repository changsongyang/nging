// Package autofill provides the Chrome DevTools Protocol
// commands, types, and events for the Autofill domain.
//
// Defines commands and events for Autofill.
//
// Generated by the cdproto-gen command.
package autofill

// Code generated by cdproto-gen. DO NOT EDIT.

import (
	"context"

	"github.com/chromedp/cdproto/cdp"
)

// TriggerParams trigger autofill on a form identified by the fieldId. If the
// field and related form cannot be autofilled, returns an error.
type TriggerParams struct {
	FieldID cdp.BackendNodeID `json:"fieldId"` // Identifies a field that serves as an anchor for autofill.
	Card    *CreditCard       `json:"card"`    // Credit card information to fill out the form. Credit card data is not saved.
}

// Trigger trigger autofill on a form identified by the fieldId. If the field
// and related form cannot be autofilled, returns an error.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Autofill#method-trigger
//
// parameters:
//
//	fieldID - Identifies a field that serves as an anchor for autofill.
//	card - Credit card information to fill out the form. Credit card data is not saved.
func Trigger(fieldID cdp.BackendNodeID, card *CreditCard) *TriggerParams {
	return &TriggerParams{
		FieldID: fieldID,
		Card:    card,
	}
}

// Do executes Autofill.trigger against the provided context.
func (p *TriggerParams) Do(ctx context.Context) (err error) {
	return cdp.Execute(ctx, CommandTrigger, p, nil)
}

// Command names.
const (
	CommandTrigger = "Autofill.trigger"
)
