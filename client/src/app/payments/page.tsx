import Link from 'next/link';

export default function PaymentsPage() {
  return (
    <div>
      <h1>Payments Page</h1>
      <Link href="/payments/custom">Custom Payments Page</Link>
      <Link href="/payments/strip">Stripe Payments Page</Link>
    </div>
  );
}
