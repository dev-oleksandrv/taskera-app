import Link from "next/link";

export const BaseHeader = () => (
  <header className="px-4 lg:px-6 h-14 flex items-center">
    <Link className="flex items-center justify-center" href="/">
      <span className="ml-2 text-xl font-bold text-teal-700">Taskera</span>
    </Link>
    <nav className="ml-auto flex gap-4 sm:gap-6">
      <Link
        className="text-sm font-medium hover:underline underline-offset-4"
        href="#"
      >
        Home
      </Link>
      <Link
        className="text-sm font-medium hover:underline underline-offset-4"
        href="#"
      >
        Features
      </Link>
      <Link
        className="text-sm font-medium hover:underline underline-offset-4"
        href="#"
      >
        Pricing
      </Link>
      <Link
        className="text-sm font-medium hover:underline underline-offset-4"
        href="/login"
      >
        Login
      </Link>
    </nav>
  </header>
);
