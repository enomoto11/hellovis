import { memo } from 'react';
import {
  Text,
  Button,
  Stack,
  TextInput,
  Box,
  Group,
  Select,
} from '@mantine/core';
import { useForm } from '@mantine/form';
import { AddStudentSchemaType } from './schema/AddStudent';

export const AddStudent = memo(() => {
  const form = useForm<AddStudentSchemaType>({
    initialValues: {
      manavisCode: '',
      firstName: '',
      lastName: '',
      grade: '',
      isHighSchool: true,
    },
  });

  return (
    // TODO: zod由来のバリデーションメッセージが表示できるようにする
    <Stack align="center" spacing="lg" h={'100%'}>
      <h1>🦄New Students🎉</h1>
      <Text fz="lg">新たに生徒を登録しましょう🤩</Text>
      <Box maw={300} mx="auto" p={40}>
        <form onSubmit={form.onSubmit((values) => console.log(values))}>
          <TextInput
            pb={20}
            withAsterisk
            required
            label="マナビス生番号"
            placeholder="000001"
            {...form.getInputProps('manavisCode')}
          />
          <TextInput
            pb={20}
            withAsterisk
            required
            label="姓"
            placeholder="マナビス"
            {...form.getInputProps('lastName')}
          />
          <TextInput
            pb={20}
            withAsterisk
            required
            label="名"
            placeholder="太郎"
            {...form.getInputProps('firstName')}
          />
          <Select
            pb={20}
            withAsterisk
            required
            label="学年"
            data={[
              { value: '1', label: '1年生' },
              { value: '2', label: '2年生' },
              { value: '3', label: '3年生' },
            ]}
            {...form.getInputProps('grade')}
          />
          <Select
            pb={20}
            withAsterisk
            required
            label="高校生/中学生"
            data={[
              { value: 'true', label: '高校生' },
              { value: 'false', label: '中学生' },
            ]}
            defaultValue={'true'}
            {...form.getInputProps('isHighSchool')}
          />

          <Group position="right" mt="md">
            <Button type="submit">登録</Button>
          </Group>
        </form>
      </Box>
    </Stack>
  );
});
