'use client';

import { useState } from 'react';

export default function CustomPaymentsPage() {
  const [success, setSuccess] = useState(false);
  return (
    <div>
      <h1>Custom Payments Page</h1>

      <button
        onClick={async () => {
          const res = await fetch('http://localhost:8080/payments', {
            method: 'POST',
            body: JSON.stringify({
              amount_usd: 500,
              customer_id: 'cus_Qjlq6Bl2Bb2nTq',
              payment_method_id: 'pm_1Pu1lRG3GvT31NaVVXODjjxN',
            }),
          });
          if (res.ok) {
            const data = await res.json();
            console.log(data);
            setSuccess(true);
          }
        }}
      >
        Make Payment
      </button>
      {success && <div>Success</div>}
    </div>
  );
}
