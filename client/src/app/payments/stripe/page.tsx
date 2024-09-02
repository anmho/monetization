'use client';
import axios from 'axios';

import { useRouter } from 'next/navigation';
import React, { useState, useEffect } from 'react';

const ProductDisplay = () => {
  const router = useRouter();
  return (
    <section>
      <div>
        <img
          src="https://i.imgur.com/EHyR2nP.png"
          alt="The cover of Stubborn Attachments"
        />
        <div>
          <h3>Stubborn Attachments</h3>
          <h5>$20.00</h5>
        </div>
      </div>
      {/* <form action="/create-checkout-session" method="POST"> */}

      <button
        type="submit"
        onClick={async () => {
          const res = await fetch('http://localhost:8080/checkout-session', {
            method: 'POST',
            body: JSON.stringify({
              customer_id: 'cus_Qjlq6Bl2Bb2nTq',
              items: [
                {
                  product_id: 'prod_QlvxMcEk08n8mi',
                  quantity: 1,
                },
              ],
            }),
          });

          const data = await res.json();
          const checkoutUrl = data['checkout_url'];

          console.log(checkoutUrl);

          if (res) {
            router.push(checkoutUrl);
          }
        }}
      >
        Checkout
      </button>
      {/* </form> */}
    </section>
  );
};

const Message = ({ message }) => (
  <section>
    <p>{message}</p>
  </section>
);

export default function StripePaymentsPage() {
  const [message, setMessage] = useState('');

  useEffect(() => {
    // Check to see if this is a redirect back from Checkout
    const query = new URLSearchParams(window.location.search);

    if (query.get('success')) {
      setMessage('Order placed! You will receive an email confirmation.');
    }

    if (query.get('canceled')) {
      setMessage(
        "Order canceled -- continue to shop around and checkout when you're ready."
      );
    }
  }, []);

  return message ? <Message message={message} /> : <ProductDisplay />;
}
