import { useEffect, useState } from "react";
import { Elements } from "@stripe/react-stripe-js";
import { loadStripe } from "@stripe/stripe-js";
import CheckoutForm from "./CheckoutForm";
import { useParams } from "react-router-dom";
import axios from "axios"; // Import Axios
function Payment() {
  const [stripePromise, setStripePromise] = useState(null);
  const [clientSecret, setClientSecret] = useState("");
  const { productId } = useParams();

  useEffect(() => {
    // Replace the fetch request with Axios
    axios
      .get("http://localhost:8080/v1/config") // Make sure to include the http://
      .then((response) => {
        // const {STRIPE_PUBLISHABLE_KEY } = response.data;
        setStripePromise(loadStripe('pk_test_51NwRJsSJH111VkZjIAOtjs6FFnvEMGKfXXVG9VwWHuTq4zN5AQ7Ils8PZRmKj6WJRtC2wsxv3r3wCBhUsOTOTBBr00YBOKSv0q'));
      })
      .catch((error) => {
        console.error("Error fetching Stripe publishable key:", error);
      });
  }, []);

  useEffect(() => {
    fetch("http://localhost:8080/v1/create-payment-intent", {
      method: "POST",
      body: JSON.stringify({
        id: parseInt(productId),
      }),
    }).then(async (result) => {
      var { clientSecret } = await result.json();
      setClientSecret(clientSecret);
    });
  }, [productId]);


  return (
    <div style={{}}>
      <h1>React Stripe and the Payment Element</h1>
      {clientSecret && stripePromise && (
        <Elements
          stripe={stripePromise}
          options={{ clientSecret, locale: "ja" }}
        >
          <CheckoutForm />
        </Elements>
      )}
    </div>
  );
}

export default Payment;