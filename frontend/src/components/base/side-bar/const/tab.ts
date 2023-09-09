import {
  IconFingerprint,
  IconUsers,
  IconReceiptRefund,
  IconFileAnalytics,
} from '@tabler/icons-react';
import {
  ATTENDANCE_LABEL,
  STUDENT_LABEL,
  EVENT_LABEL,
  HISTORY_LABEL,
} from './label';

export const ACCOUNT_TAB = 'account';
export const GENERAL_TAB = 'general';

export const tabs = {
  account: [
    {
      link: '',
      label: ATTENDANCE_LABEL,
      icon: IconFingerprint,
      onClick: () => {
        window.location.href = '/attendance';
      },
    },
  ],
  general: [
    {
      link: '/students',
      label: STUDENT_LABEL,
      icon: IconUsers,
      onClick: () => {
        window.location.href = '/students';
      },
    },
    {
      link: '/events',
      label: EVENT_LABEL,
      icon: IconReceiptRefund,
      onClick: () => {
        window.location.href = '/events';
      },
    },
    {
      link: '/histories',
      label: HISTORY_LABEL,
      icon: IconFileAnalytics,
      onClick: () => {
        window.location.href = '/histories';
      },
    },
  ],
};
