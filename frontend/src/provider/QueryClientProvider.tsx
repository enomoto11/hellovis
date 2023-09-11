import { memo } from 'react';
import { QueryClient } from 'react-query';
import { QueryClientProvider as Provider } from 'react-query';

interface Props {
  children: JSX.Element;
}

export const QueryClientProvider = memo((props: Props) => {
  const queryClient = new QueryClient();
  return <Provider client={queryClient}>{props.children || <></>}</Provider>;
});
