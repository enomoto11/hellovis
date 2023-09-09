import { memo } from 'react';
import {
  Navbar,
  SegmentedControl,
  createStyles,
  getStylesRef,
  rem,
} from '@mantine/core';
import { IconLogout } from '@tabler/icons-react';
import { tabs } from './const/tab';
import { useSideBar } from './useSideBar';

export const SideBar = memo(() => {
  const { classes, cx } = useStyles();
  const {
    section,
    active,
    setActive,
    logout,
    segmentedControlData,
    segmentedControlOnChange,
  } = useSideBar();

  const links = tabs[section].map((tab) => (
    <a
      className={cx(classes.link, {
        [classes.linkActive]: tab.label === active,
      })}
      href={tab.link}
      key={tab.label}
      onClick={(event) => {
        event.preventDefault();
        setActive(tab.label);
        tab.onClick();
      }}
    >
      <tab.icon className={classes.linkIcon} stroke={1.5} />
      <span>{tab.label}</span>
    </a>
  ));

  return (
    <Navbar width={{ sm: 300 }} p="md" className={classes.navbar}>
      <Navbar.Section>
        <SegmentedControl
          value={section}
          onChange={segmentedControlOnChange}
          transitionTimingFunction="ease"
          fullWidth
          data={segmentedControlData}
        />
      </Navbar.Section>

      <Navbar.Section grow mt="xl">
        {links}
      </Navbar.Section>

      <Navbar.Section className={classes.footer}>
        <a className={classes.link} onClick={() => logout()}>
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
