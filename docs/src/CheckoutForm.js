import { PaymentElement } from "@stripe/react-stripe-js";
import { useState } from "react";
import { useStripe, useElements } from "@stripe/react-stripe-js";
import { ElementsConsumer } from '@stripe/react-stripe-js';

export default function CheckoutForm() {
  const stripe = useStripe();
  const elements = useElements();

  const [message, setMessage] = useState(null);
  const [isProcessing, setIsProcessing] = useState(false);

  const handleSubmit = async (e) => {
    e.preventDefault();

    if (!stripe || !elements) {
      // Stripe.js has not fully loaded yet.
      // Please wait and try again later.
      setMessage("Stripe.js is not fully loaded. Please try again later.");
      return;
    }

    setIsProcessing(true);

    const { error } = await stripe.confirmPayment({
      elements,
      confirmParams: {
        return_url: `${window.location.origin}/completion-us`,
      },
    });

    if (error.type === "card_error" || error.type === "validation_error") {
      // There was an issue with your payment or card information.
      setMessage(error.message);
    } else {
      // An unexpected error occurred. Please try again later.
      setMessage("An unexpected error occurred.");
    }

    setIsProcessing(false);
  };

  return (
    <ElementsConsumer>
      {({ stripe, elements }) => (
        <form id="payment-form" onSubmit={handleSubmit}>
          <label htmlFor="payment-element">Card Details</label>
          <PaymentElement id="payment-element" />
          <button
            disabled={isProcessing || !stripe || !elements}
            id="submit"
            type="submit"
          >
            <span id="button-text">
              {isProcessing ? "Processing ... " : "Pay Now"}
            </span>
          </button>
          {/* Display any error or success messages */}
          {message && <div id="payment-message">{message}</div>}
        </form>
      )}
    </ElementsConsumer>
  );
}
