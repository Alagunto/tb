package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/alagunto/tb"
	"github.com/alagunto/tb/request"
	"github.com/alagunto/tb/telegram"
)

// Payment provider token - get from @BotFather using /setpaymentprovider
// For testing, you can use Stripe test token
var paymentProviderToken = os.Getenv("PAYMENT_PROVIDER_TOKEN")

func main() {
	token := os.Getenv("BOT_TOKEN")
	if token == "" {
		log.Fatal("BOT_TOKEN environment variable is required")
	}

	if paymentProviderToken == "" {
		log.Println("⚠️  Warning: PAYMENT_PROVIDER_TOKEN not set")
		log.Println("To test payments, get a provider token from @BotFather:")
		log.Println("1. Send /setpaymentprovider to @BotFather")
		log.Println("2. Choose your bot")
		log.Println("3. Select a payment provider (Stripe recommended for testing)")
		log.Println("4. Set the token as environment variable")
	}

	// Create a request builder function
	requestBuilder := func(req request.Interface) (*request.Native, error) {
		return request.NewNativeFromRequest(req), nil
	}

	// Create bot settings - include payment updates
	settings := tb.Settings[*request.Native]{
		Token: token,
		Poller: &tb.LongPoller{
			Timeout:        10 * time.Second,
			AllowedUpdates: []string{"message", "pre_checkout_query", "shipping_query"},
		},
		OnError: func(err error, ctx *request.Native) {
			log.Printf("Error: %v", err)
		},
	}

	// Create the bot
	bot, err := tb.NewBot(requestBuilder, settings)
	if err != nil {
		log.Fatalf("Failed to create bot: %v", err)
	}

	// Start command
	bot.Handle("/start", func(c *request.Native) error {
		return c.Reply("💳 Welcome to Payments Demo!\n\n" +
			"This bot demonstrates Telegram Payments API.\n\n" +
			"Features:\n" +
			"• Send invoices\n" +
			"• Accept payments\n" +
			"• Shipping options\n" +
			"• Tips support\n" +
			"• Order validation\n\n" +
			"Commands:\n" +
			"/buy_digital - Buy digital product (no shipping)\n" +
			"/buy_physical - Buy physical product (with shipping)\n" +
			"/buy_subscription - Monthly subscription\n" +
			"/donate - Donation with tip\n\n" +
			"⚠️ Note: Requires payment provider token from @BotFather\n" +
			"Test payments don't charge real money!")
	})

	// Buy digital product (no shipping)
	bot.Handle("/buy_digital", func(c *request.Native) error {
		if paymentProviderToken == "" {
			return c.Reply("❌ Payment provider token not configured. See bot startup message.")
		}

		// Create invoice
		invoice := &telegram.Invoice{
			Title:       "Premium Bot Access",
			Description: "Get premium features for 1 month:\n• No ads\n• Priority support\n• Exclusive content",
			Payload:     "premium_monthly_001", // Internal identifier
			ProviderToken: paymentProviderToken,
			Currency:    "USD",
			Prices: []telegram.LabeledPrice{
				{Label: "Premium Access", Amount: 990}, // $9.90 (amount in cents)
				{Label: "Processing Fee", Amount: 10},  // $0.10
			},
			PhotoURL:    "https://via.placeholder.com/400x300/4A90E2/ffffff?text=Premium",
			PhotoWidth:  400,
			PhotoHeight: 300,
			NeedName:    false, // Don't need user's name
			NeedPhoneNumber: false,
			NeedEmail:   true, // Need email for digital delivery
			NeedShippingAddress: false,
			SendPhoneNumberToProvider: false,
			SendEmailToProvider: true,
			IsFlexible: false, // Fixed price
		}

		if err := c.SendInvoice(invoice); err != nil {
			return c.Reply("Failed to create invoice: " + err.Error())
		}

		return nil
	})

	// Buy physical product (with shipping)
	bot.Handle("/buy_physical", func(c *request.Native) error {
		if paymentProviderToken == "" {
			return c.Reply("❌ Payment provider token not configured. See bot startup message.")
		}

		invoice := &telegram.Invoice{
			Title:       "Bot Merchandise T-Shirt",
			Description: "Official bot t-shirt\n• 100% Cotton\n• Available in all sizes\n• Free sticker included",
			Payload:     "tshirt_blue_xl_001",
			ProviderToken: paymentProviderToken,
			Currency:    "USD",
			Prices: []telegram.LabeledPrice{
				{Label: "T-Shirt", Amount: 2999}, // $29.99
			},
			PhotoURL:    "https://via.placeholder.com/400x400/FF6B6B/ffffff?text=T-Shirt",
			PhotoWidth:  400,
			PhotoHeight: 400,
			NeedName:    true,
			NeedPhoneNumber: true,
			NeedEmail:   true,
			NeedShippingAddress: true, // Required for physical goods
			IsFlexible: true, // Price varies by shipping option
		}

		if err := c.SendInvoice(invoice); err != nil {
			return c.Reply("Failed to create invoice: " + err.Error())
		}

		return nil
	})

	// Subscription
	bot.Handle("/buy_subscription", func(c *request.Native) error {
		if paymentProviderToken == "" {
			return c.Reply("❌ Payment provider token not configured. See bot startup message.")
		}

		invoice := &telegram.Invoice{
			Title:       "Pro Subscription",
			Description: "Unlock all features:\n• Unlimited requests\n• Priority queue\n• Advanced analytics\n• API access",
			Payload:     "subscription_pro_monthly",
			ProviderToken: paymentProviderToken,
			Currency:    "USD",
			Prices: []telegram.LabeledPrice{
				{Label: "Pro Subscription (1 month)", Amount: 1999}, // $19.99
			},
			PhotoURL:    "https://via.placeholder.com/400x300/9B59B6/ffffff?text=Pro",
			PhotoWidth:  400,
			PhotoHeight: 300,
			NeedEmail:   true,
		}

		if err := c.SendInvoice(invoice); err != nil {
			return c.Reply("Failed to create invoice: " + err.Error())
		}

		return nil
	})

	// Donation with tip
	bot.Handle("/donate", func(c *request.Native) error {
		if paymentProviderToken == "" {
			return c.Reply("❌ Payment provider token not configured. See bot startup message.")
		}

		invoice := &telegram.Invoice{
			Title:       "Support the Bot",
			Description: "Help us keep the bot running and add new features!\n\nYour support means a lot ❤️",
			Payload:     "donation_001",
			ProviderToken: paymentProviderToken,
			Currency:    "USD",
			Prices: []telegram.LabeledPrice{
				{Label: "Donation", Amount: 500}, // $5.00 base
			},
			PhotoURL:    "https://via.placeholder.com/400x300/F39C12/ffffff?text=Donate",
			PhotoWidth:  400,
			PhotoHeight: 300,
			// Suggested tip amounts
			SuggestedTipAmounts: []int{100, 300, 500, 1000}, // $1, $3, $5, $10
			MaxTipAmount: 10000, // Max $100
		}

		if err := c.SendInvoice(invoice); err != nil {
			return c.Reply("Failed to create invoice: " + err.Error())
		}

		return nil
	})

	// Handle shipping queries (for physical products)
	bot.Handle(tb.OnShipping, func(c *request.Native) error {
		query := c.ShippingQuery()
		if query == nil {
			return nil
		}

		log.Printf("Shipping query from %s to %s",
			query.Sender.FirstName,
			query.ShippingAddress.CountryCode,
		)

		// Define shipping options based on location
		var options []telegram.ShippingOption

		// Domestic shipping
		if query.ShippingAddress.CountryCode == "US" {
			options = append(options, telegram.ShippingOption{
				ID:    "us_standard",
				Title: "Standard Shipping (5-7 days)",
				Prices: []telegram.LabeledPrice{
					{Label: "Standard", Amount: 500}, // $5.00
				},
			})

			options = append(options, telegram.ShippingOption{
				ID:    "us_express",
				Title: "Express Shipping (2-3 days)",
				Prices: []telegram.LabeledPrice{
					{Label: "Express", Amount: 1500}, // $15.00
				},
			})
		} else {
			// International shipping
			options = append(options, telegram.ShippingOption{
				ID:    "intl_standard",
				Title: "International Shipping (10-15 days)",
				Prices: []telegram.LabeledPrice{
					{Label: "International", Amount: 2000}, // $20.00
				},
			})
		}

		// Answer shipping query with options
		return c.API.AnswerShippingQuery(query, options, true, "")
	})

	// Handle pre-checkout queries (final validation before payment)
	bot.Handle(tb.OnPreCheckout, func(c *request.Native) error {
		query := c.PreCheckoutQuery()
		if query == nil {
			return nil
		}

		log.Printf("Pre-checkout query from %s: %s (%d %s)",
			query.Sender.FirstName,
			query.InvoicePayload,
			query.TotalAmount,
			query.Currency,
		)

		// Validate the order
		// In a real app, you would:
		// 1. Check inventory
		// 2. Validate prices
		// 3. Check user eligibility
		// 4. Verify payment details

		// Example validation
		if query.TotalAmount > 100000 { // > $1000
			// Reject the payment
			return c.API.AnswerPreCheckoutQuery(query, false, "Amount too large. Contact support.")
		}

		// Accept the payment
		return c.API.AnswerPreCheckoutQuery(query, true, "")
	})

	// Handle successful payments
	bot.Handle(tb.OnSuccessfulPayment, func(c *request.Native) error {
		msg := c.Message()
		if msg == nil || msg.SuccessfulPayment == nil {
			return nil
		}

		payment := msg.SuccessfulPayment

		log.Printf("Payment successful! ID: %s, Amount: %d %s",
			payment.TelegramPaymentChargeID,
			payment.TotalAmount,
			payment.Currency,
		)

		// Process the order
		// In a real app, you would:
		// 1. Save payment to database
		// 2. Grant access/ship product
		// 3. Send receipt
		// 4. Notify admins

		confirmation := fmt.Sprintf("✅ <b>Payment Successful!</b>\n\n"+
			"Order ID: <code>%s</code>\n"+
			"Amount: %.2f %s\n"+
			"Invoice Payload: %s\n\n"+
			"Thank you for your purchase! 🎉\n\n"+
			"You will receive a confirmation email shortly.",
			payment.TelegramPaymentChargeID,
			float64(payment.TotalAmount)/100,
			payment.Currency,
			payment.InvoicePayload,
		)

		return c.Reply(confirmation, tb.Send().WithParseMode(telegram.ParseModeHTML))
	})

	// Handle text messages
	bot.Handle(tb.OnText, func(c *request.Native) error {
		msg := c.Message()
		if msg == nil || msg.Text == "" || msg.Text[0] == '/' {
			return nil
		}

		return c.Reply("Send /start to see available products!")
	})

	log.Println("💳 Payments bot started! Press Ctrl+C to stop.")
	if paymentProviderToken != "" {
		log.Println("✅ Payment provider configured")
	} else {
		log.Println("⚠️  Payment provider NOT configured - get token from @BotFather")
	}
	log.Println("💡 Test with /buy_digital for digital products (no shipping)")
	log.Println("📦 Test with /buy_physical for physical products (with shipping)")
	bot.Start()
}

