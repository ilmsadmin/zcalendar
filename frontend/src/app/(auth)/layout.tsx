export default function AuthLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <div className="flex min-h-[100dvh] items-center justify-center bg-gradient-to-br from-amber-50 via-white to-red-50 dark:from-gray-950 dark:via-gray-900 dark:to-gray-950 px-4 py-8">
      <div className="w-full max-w-sm sm:max-w-md">{children}</div>
    </div>
  );
}
