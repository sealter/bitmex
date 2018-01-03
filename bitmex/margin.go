/*
 * BitMEX API
 *
 * ## REST API for the BitMEX Trading Platform  [Changelog](/app/apiChangelog)    #### Getting Started   ##### Fetching Data  All REST endpoints are documented below. You can try out any query right from this interface.  Most table queries accept `count`, `start`, and `reverse` params. Set `reverse=true` to get rows newest-first.  Additional documentation regarding filters, timestamps, and authentication is available in [the main API documentation](https://www.bitmex.com/app/restAPI).  *All* table data is available via the [Websocket](/app/wsAPI). We highly recommend using the socket if you want to have the quickest possible data without being subject to ratelimits.  ##### Return Types  By default, all data is returned as JSON. Send `?_format=csv` to get CSV data or `?_format=xml` to get XML data.  ##### Trade Data Queries  *This is only a small subset of what is available, to get you started.*  Fill in the parameters and click the `Try it out!` button to try any of these queries.  * [Pricing Data](#!/Quote/Quote_get)  * [Trade Data](#!/Trade/Trade_get)  * [OrderBook Data](#!/OrderBook/OrderBook_getL2)  * [Settlement Data](#!/Settlement/Settlement_get)  * [Exchange Statistics](#!/Stats/Stats_history)  Every function of the BitMEX.com platform is exposed here and documented. Many more functions are available.  -  ## All API Endpoints  Click to expand a section.
 *
 * OpenAPI spec version: 1.2.0
 * Contact: support@bitmex.com
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package bitmex

import (
	"time"
)

type Margin struct {
	Account float32 `json:"account,omitempty"`

	Currency string `json:"currency,omitempty"`

	RiskLimit float32 `json:"riskLimit,omitempty"`

	PrevState string `json:"prevState,omitempty"`

	State string `json:"state,omitempty"`

	Action string `json:"action,omitempty"`

	Amount float32 `json:"amount,omitempty"`

	PendingCredit float32 `json:"pendingCredit,omitempty"`

	PendingDebit float32 `json:"pendingDebit,omitempty"`

	ConfirmedDebit float32 `json:"confirmedDebit,omitempty"`

	PrevRealisedPnl float32 `json:"prevRealisedPnl,omitempty"`

	PrevUnrealisedPnl float32 `json:"prevUnrealisedPnl,omitempty"`

	GrossComm float32 `json:"grossComm,omitempty"`

	GrossOpenCost float32 `json:"grossOpenCost,omitempty"`

	GrossOpenPremium float32 `json:"grossOpenPremium,omitempty"`

	GrossExecCost float32 `json:"grossExecCost,omitempty"`

	GrossMarkValue float32 `json:"grossMarkValue,omitempty"`

	RiskValue float32 `json:"riskValue,omitempty"`

	TaxableMargin float32 `json:"taxableMargin,omitempty"`

	InitMargin float32 `json:"initMargin,omitempty"`

	MaintMargin float32 `json:"maintMargin,omitempty"`

	SessionMargin float32 `json:"sessionMargin,omitempty"`

	TargetExcessMargin float32 `json:"targetExcessMargin,omitempty"`

	VarMargin float32 `json:"varMargin,omitempty"`

	RealisedPnl float32 `json:"realisedPnl,omitempty"`

	UnrealisedPnl float32 `json:"unrealisedPnl,omitempty"`

	IndicativeTax float32 `json:"indicativeTax,omitempty"`

	UnrealisedProfit float32 `json:"unrealisedProfit,omitempty"`

	SyntheticMargin float32 `json:"syntheticMargin,omitempty"`

	WalletBalance float32 `json:"walletBalance,omitempty"`

	MarginBalance float32 `json:"marginBalance,omitempty"`

	MarginBalancePcnt float64 `json:"marginBalancePcnt,omitempty"`

	MarginLeverage float64 `json:"marginLeverage,omitempty"`

	MarginUsedPcnt float64 `json:"marginUsedPcnt,omitempty"`

	ExcessMargin float32 `json:"excessMargin,omitempty"`

	ExcessMarginPcnt float64 `json:"excessMarginPcnt,omitempty"`

	AvailableMargin float32 `json:"availableMargin,omitempty"`

	WithdrawableMargin float32 `json:"withdrawableMargin,omitempty"`

	Timestamp time.Time `json:"timestamp,omitempty"`

	GrossLastValue float32 `json:"grossLastValue,omitempty"`

	Commission float64 `json:"commission,omitempty"`
}
