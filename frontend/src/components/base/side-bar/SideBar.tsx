import { memo, useCallback, useState } from 'react';
import {
  Navbar,
  SegmentedControl,
  createStyles,
  getStylesRef,
  rem,
} from '@mantine/core';
import {
  IconFingerprint,
  IconUsers,
  IconFileAnalytics,
  IconReceiptRefund,
  IconLogout,
} from '@tabler/icons-react';
import { useAuth0 } from '@auth0/auth0-react';

const tabs = {
  account: [{ link: '', label: '登下校を管理', icon: IconFingerprint }],
  general: [
    { link: '', label: '生徒管理', icon: IconUsers },
    { link: '', label: '行事管理', icon: IconReceiptRefund },
    { link: '', label: '来校記録', icon: IconFileAnalytics },
  ],
};

// TODO: buisiness logic と DOM の分離を行う
export const NavbarSegmented = memo(() => {
  const { classes, cx } = useStyles();
  const [section, setSection] = useState<'account' | 'general'>('account');
  const [active, setActive] = useState('登下校を管理');

  const links = tabs[section].map((item) => (
    <a
      className={cx(classes.link, {
        [classes.linkActive]: item.label === active,
      })}
      href={item.link}
      key={item.label}
      onClick={(event) => {
        event.preventDefault();
        setActive(item.label);
      }}
    >
      <item.icon className={classes.linkIcon} stroke={1.5} />
      <span>{item.label}</span>
    </a>
  ));

  const { logout } = useAuth0();

  const handleLoguout = useCallback(async () => {
    await logout();
  }, [logout]);

  return (
    <Navbar width={{ sm: 300 }} p="md" className={classes.navbar}>
      <Navbar.Section>
        <SegmentedControl
          value={section}
          onChange={(value: 'account' | 'general') => setSection(value)}
          transitionTimingFunction="ease"
          fullWidth
          data={[
            { label: 'マナビス生', value: 'account' },
            { label: 'Administrator', value: 'general' },
          ]}
        />
      </Navbar.Section>

      <Navbar.Section grow mt="xl">
        {links}
      </Navbar.Section>

      <Navbar.Section className={classes.footer}>
        <a className={classes.link} onClick={handleLoguout}>
          <IconLogout className={classes.linkIcon} stroke={1.5} />
          <span>システムを終了する</span>
        </a>
      </Navbar.Section>
    </Navbar>
  );
});

const useStyles = createStyles((theme) => ({
  navbar: {
    backgroundColor:
      theme.colorScheme === 'dark' ? theme.colors.dark[7] : theme.white,
  },

  title: {
    textTransform: 'uppercase',
    letterSpacing: rem(-0.25),
  },

  link: {
    ...theme.fn.focusStyles(),
    display: 'flex',
    alignItems: 'center',
    textDecoration: 'none',
    fontSize: theme.fontSizes.sm,
    color:
      theme.colorScheme === 'dark'
        ? theme.colors.dark[1]
        : theme.colors.gray[7],
    padding: `${theme.spacing.xs} ${theme.spacing.sm}`,
    borderRadius: theme.radius.sm,
    fontWeight: 500,

    '&:hover': {
      backgroundColor:
        theme.colorScheme === 'dark'
          ? theme.colors.dark[6]
          : theme.colors.gray[0],
      color: theme.colorScheme === 'dark' ? theme.white : theme.black,

      [`& .${getStylesRef('icon')}`]: {
        color: theme.colorScheme === 'dark' ? theme.white : theme.black,
      },
    },
  },

  linkIcon: {
    ref: getStylesRef('icon'),
    color:
      theme.colorScheme === 'dark'
        ? theme.colors.dark[2]
        : theme.colors.gray[6],
    marginRight: theme.spacing.sm,
  },

  linkActive: {
    '&, &:hover': {
      backgroundColor: theme.fn.variant({
        variant: 'light',
        color: theme.primaryColor,
      }).background,
      color: theme.fn.variant({ variant: 'light', color: theme.primaryColor })
        .color,
      [`& .${getStylesRef('icon')}`]: {
        color: theme.fn.variant({ variant: 'light', color: theme.primaryColor })
          .color,
      },
    },
  },

  footer: {
    borderTop: `${rem(1)} solid ${
      theme.colorScheme === 'dark' ? theme.colors.dark[4] : theme.colors.gray[3]
    }`,
    paddingTop: theme.spacing.md,
  },
}));
