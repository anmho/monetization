'use client';
import { Stripe } from '@stripe/stripe-js';
import { useEffect, useState } from 'react';

interface CompletionPageProps {
  stripePromise: Promise<Stripe>;
}

function CompletionPage(props: CompletionPageProps) {
  const [messageBody, setMessageBody] = useState('');
  const { stripePromise } = props;

  useEffect(() => {
    if (!stripePromise) return;

    stripePromise.then(async (stripe) => {
      const url = new URL(window.location.toString());
      const clientSecret = url.searchParams.get(
        'payment_intent_client_secret'
      ) as string;
      const { error, paymentIntent } = await stripe.retrievePaymentIntent(
        clientSecret
      );

      setMessageBody(
        error
          ? `> ${error.message}`
          : `> Payment ${paymentIntent.status}: <a href="https://dashboard.stripe.com/test/payments/${paymentIntent.id}" target="_blank" rel="noreferrer">${paymentIntent.id}</a>`
      );
    });
  }, [stripePromise]);

  return (
    <>
      <h1>Thank you!</h1>
      <a href="/">home</a>
      <div
        id="messages"
        role="alert"
        style={messageBody ? { display: 'block' } : {}}
      >
        {messageBody}
      </div>
    </>
  );
}

export default CompletionPage;
