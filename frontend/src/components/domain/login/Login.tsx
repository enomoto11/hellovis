import { PopupConfigOptions, PopupLoginOptions } from '@auth0/auth0-react';
import {
  Paper,
  createStyles,
  Button,
  Title,
  MantineTheme,
} from '@mantine/core';
import { memo, useCallback } from 'react';

interface Props {
  loginWithPopup: (
    // eslint-disable-next-line no-unused-vars
    options?: PopupLoginOptions | undefined,
    // eslint-disable-next-line no-unused-vars
    config?: PopupConfigOptions | undefined,
  ) => Promise<void>;
}

export const Login = memo((props: Props) => {
  const { classes } = useStyles();

  const params = useLogin(props);

  return (
    <Paper className={classes.paper} radius="lg" p={30}>
      <Title order={2} className={classes.title} align="center" mt="md" mb={50}>
        This is Hellovis.
      </Title>

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

const useLogin = (props: Props) => {
  const { loginWithPopup } = props;

  const onClick = useCallback(() => {
    loginWithPopup();
  }, [loginWithPopup]);

  return {
    onClick,
  };
};
