import { memo } from 'react';
import { AddStudent } from './AddStudent/AddStudent';
import { Flex } from '@mantine/core';

export const Students = memo(() => {
  return (
    <Flex
      direction={{ base: 'column', sm: 'row' }}
      gap={{ base: 'sm', sm: 'lg' }}
      justify={{ sm: 'center' }}
    >
      <AddStudent />
    </Flex>
  );
});
