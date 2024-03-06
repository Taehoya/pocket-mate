import { QueryClient } from '@tanstack/react-query';
import { cache } from 'react';

const getQueryClient = (() => new QueryClient());
export default getQueryClient;