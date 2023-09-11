import { memo } from 'react';
import { Button } from '@mantine/core';
import { useAuth0 } from '@auth0/auth0-react';

export const Home = memo(() => {
  const authParams = useAuth0();

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
