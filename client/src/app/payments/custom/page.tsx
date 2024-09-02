'use client';

import { PaymentElement } from '@stripe/react-stripe-js';
import { useEffect, useState } from 'react';
import Payment from './payment';

export default function CustomPaymentsPage() {
  const [isLoading, setIsLoading] = useState(true);
  const [clientSecret, setClientSecret] = useState('');

  const paymentParams = {
    amount_usd: 500,
    customer_id: 'cus_Qjlq6Bl2Bb2nTq',
    payment_method_id: 'pm_1Pu1lRG3GvT31NaVVXODjjxN',
  };

  // create a payment intent on page load
  // try using rsc maybe

  useEffect(() => {
    const createPaymentIntent = async () => {
      const res = await fetch('http://localhost:8080/payments', {
        method: 'POST',
        body: JSON.stringify(paymentParams),
      });
      const data = await res.json();
      const { client_secret } = data;

      console.log(data);
      setClientSecret(client_secret);
      setIsLoading(false);
      // .then((res) => res.json())
      // .then((data) => {
      //   setClientSecret(data['client_secret']);
      // });
    };

    createPaymentIntent();
  }, []);
  return (
    <div>
      <h1>Custom Payments Page</h1>
      {isLoading ? <p>Loading</p> : <Payment clientSecret={clientSecret} />}
    </div>
  );
}
