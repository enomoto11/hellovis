import { useState, useMemo, useCallback } from 'react';
import { ATTENDANCE_LABEL } from './const/label';
import { ACCOUNT_TAB, GENERAL_TAB } from './const/tab';
import { useAuth } from '../../../provider/Auth0provider';

export const useSideBar = () => {
  const [section, setSection] = useState<
    typeof ACCOUNT_TAB | typeof GENERAL_TAB
  >(ACCOUNT_TAB);
  const [active, setActive] = useState(ATTENDANCE_LABEL);

  const authParams = useAuth();

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

  const logout = useCallback(() => {
    authParams?.logout();
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
