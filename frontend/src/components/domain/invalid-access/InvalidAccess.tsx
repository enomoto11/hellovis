import { Paper, createStyles, Title, MantineTheme } from '@mantine/core';
import { memo } from 'react';

export const InvalidAccess = memo(() => {
  const { classes } = useStyles();

  return (
    <Paper className={classes.paper} radius="lg" p={30}>
      <Title order={2} className={classes.title} align="center" mt="md" mb={50}>
        You have no right to access this page.
      </Title>
    </Paper>
  );
});

const useStyles = createStyles((theme: MantineTheme) => ({
  paper: {
    alignContent: 'center',
    width: 'calc(100% / 2)',
    height: 'calc(100% / 3)',
    minHeight: 600,
    maxWidth: 450,
  },

  title: {
    color: theme.colorScheme === 'dark' ? theme.white : theme.black,
    fontFamily: `Greycliff CF, ${theme.fontFamily}`,
  },
}));
