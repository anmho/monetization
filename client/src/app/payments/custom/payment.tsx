'use client';
import { Elements, PaymentElement } from '@stripe/react-stripe-js';
import { loadStripe, StripeElementsOptions } from '@stripe/stripe-js';
import CheckoutForm from './checkout-form';

const stripePromise = loadStripe(
  'pk_test_51NWbEnG3GvT31NaVIQqD5PrgevW9GBH5e8UAy2sFzMcSCVWXgDyEBA3eZfFPUriP7G3HI3y45FW8rCBi5xWp1T0800NYx01H9c'
);

export default function Payment({ clientSecret }: { clientSecret: string }) {
  const options = {
    clientSecret: clientSecret,
    // Fully customizable with appearance API.
  };
  return (
    <Elements stripe={stripePromise} options={options}>
      <CheckoutForm />
    </Elements>
  );
}
