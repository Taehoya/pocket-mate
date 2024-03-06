import ReactQueryProvider from "@/lib/react-query/ReactQueryProvider";
import { NextIntlClientProvider } from "next-intl";
import { notFound } from "next/navigation";


export function generateStaticParams() {
  return [{ locale: 'en' }, { locale: 'kr' }];
}

export default async function LocaleLayout({
  children,
  params: { locale },
}: any) {
  let messages;
  try {
    messages = (await import(`../../messages/${locale}.json`)).default;
  } catch (error) {
    notFound();
  }

  return (
    <html lang={locale}>
      <body style={{ margin: 0 }}>
        <NextIntlClientProvider locale={locale} messages={messages}>
          <ReactQueryProvider>
            {children}
          </ReactQueryProvider>
        </NextIntlClientProvider>
      </body>
    </html>
  );
}
