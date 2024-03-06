'use client';

import { createSyncStoragePersister } from '@tanstack/query-sync-storage-persister';
import { QueryClient } from '@tanstack/react-query';
import { ReactQueryDevtools } from '@tanstack/react-query-devtools';
import { PersistQueryClientProvider } from '@tanstack/react-query-persist-client';
import React, { useState } from 'react';

export const queryClient = new QueryClient({
  defaultOptions: {
    queries: {
      staleTime: Infinity,
      retry: false,
      cacheTime: 1000 * 60 * 60 * 24,
      refetchOnWindowFocus: false,
      retryDelay: 1000,
    },
  },
});

function ReactQueryProvider({ children }: React.PropsWithChildren) {
  const [client] = useState(
    () =>
      new QueryClient({
        defaultOptions: {
          queries: {
            staleTime: 3 * 60 * 1000,
            retry: 2,
            cacheTime: 3 * 60 * 1000,
            refetchOnWindowFocus: false,
            retryDelay: 1000,
          },
        },
      })
  );

  const persister = createSyncStoragePersister({
    storage: typeof window !== 'undefined' ? window.localStorage : undefined,
  });

  return (
    <PersistQueryClientProvider client={queryClient} persistOptions={{ persister }}>
      {children}
      <ReactQueryDevtools initialIsOpen={false} />
    </PersistQueryClientProvider>
  );
}

export default ReactQueryProvider;
