'use client';
import { Elements } from '@stripe/react-stripe-js';
import { loadStripe, StripeElementsOptions } from '@stripe/stripe-js';

const stripePromise = loadStripe(
  'pk_test_51NWbEnG3GvT31NaVIQqD5PrgevW9GBH5e8UAy2sFzMcSCVWXgDyEBA3eZfFPUriP7G3HI3y45FW8rCBi5xWp1T0800NYx01H9c'
);

export default function StripeProvider({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  const options: StripeElementsOptions = {
    mode: 'setup',
    currency: 'usd',
    // Fully customizable with appearance API.
    appearance: {},
  };
  return (
    <Elements stripe={stripePromise} options={options}>
      {children}
    </Elements>
  );
}
