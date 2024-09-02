import Link from 'next/link';

export default function SubscriptionsPage() {
  return (
    <div>
      <h1>Subscriptions Page</h1>
      <Link href="/subscriptions/custom">Custom Page</Link>

      <Link href="/subscriptions/stripe">Stripe Page</Link>
    </div>
  );
}
