import { useAuth0 } from '@auth0/auth0-react';
import { useState, useMemo, useCallback } from 'react';
import { ATTENDANCE_LABEL } from './const/label';
import { ACCOUNT_TAB, GENERAL_TAB } from './const/tab';

export const useSideBar = () => {
  const [section, setSection] = useState<
    typeof ACCOUNT_TAB | typeof GENERAL_TAB
  >(ACCOUNT_TAB);
  const [active, setActive] = useState(ATTENDANCE_LABEL);

  const { logout } = useAuth0();

  const segmentedControlData = useMemo(
    () => [
      { label: 'マナビス生', value: ACCOUNT_TAB },
      { label: 'Administrator', value: GENERAL_TAB },
    ],
    [],
  );
  const segmentedControlOnChange = useCallback(
    (value: typeof ACCOUNT_TAB | typeof GENERAL_TAB) => setSection(value),
    [],
  );

  const handleToStudentsPage = useCallback(() => {
    window.location.href = '/students';
  }, []);

  return {
    section,
    active,
    setActive,
    logout,
    segmentedControlData,
    segmentedControlOnChange,
    handleToStudentsPage,
  };
};
