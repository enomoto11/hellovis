import { memo } from 'react';
import { Button } from '@mantine/core';
import { useAuth } from '../../../provider/Auth0provider';

export const Home = memo(() => {
  const authParams = useAuth();

  return (
    <div>
      <h1>Home</h1>
      <Button
        onClick={() => {
          authParams?.logout();
        }}
      >
        Log out
      </Button>
    </div>
  );
});
