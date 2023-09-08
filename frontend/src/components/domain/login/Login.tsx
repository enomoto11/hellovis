import { useAuth0 } from '@auth0/auth0-react';
import {
  Paper,
  createStyles,
  Button,
  Title,
  MantineTheme,
} from '@mantine/core';
import { memo, useCallback } from 'react';
import { DOMAIN_DEV, DOMAIN_PROD } from '../../../env/auth0.env';

export const Login = memo(() => {
  const { classes } = useStyles();

  const params = useLogin();

  return (
    <Paper className={classes.paper} radius="lg" p={30}>
      <Title order={2} className={classes.title} align="center" mt="md" mb={50}>
        This is Hellovis.
      </Title>

      <h1>domain</h1>

      <h1>{DOMAIN_DEV || DOMAIN_PROD}</h1>

      <Button fullWidth mt="xl" radius="lg" size="xl" {...params}>
        Login
      </Button>
    </Paper>
  );
});

const useStyles = createStyles((theme: MantineTheme) => ({
  paper: {
    alignContent: 'center',
    width: 'calc(100% / 2)',
    height: 'calc(100% / 3 *2)',
    minHeight: 600,
    maxWidth: 450,
  },

  title: {
    color: theme.colorScheme === 'dark' ? theme.white : theme.black,
    fontFamily: `Greycliff CF, ${theme.fontFamily}`,
  },
}));

const useLogin = () => {
  const { loginWithRedirect } = useAuth0();

  const onClick = useCallback(async () => {
    await loginWithRedirect({
      appState: {
        returnTo: '/',
      },
    });
  }, [loginWithRedirect]);

  return {
    onClick,
  };
};
