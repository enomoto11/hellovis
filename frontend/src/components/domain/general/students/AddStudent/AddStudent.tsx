import { memo, useCallback, useEffect, useState } from 'react';
import {
  Text,
  Button,
  Stack,
  TextInput,
  Box,
  Group,
  Select,
  Checkbox,
} from '@mantine/core';
import { useForm } from '@mantine/form';
import { AddStudentSchemaType, addStudentSchema } from './schema/AddStudent';
import { useMutation } from 'react-query';
import axios from 'axios';

export const AddStudent = memo(() => {
  const form = useForm<AddStudentSchemaType>({
    initialValues: {
      manavisCode: '200828',
      firstName: '東悟',
      lastName: '榎本',
      grade: '2',
      isHighSchool: true,
    },
  });

  const [checked, setChecked] = useState(true);

  const handleChangeCheckbox = useCallback(() => {
    setChecked(!checked);
  }, [checked]);

  useEffect(() => {
    form.setValues({
      isHighSchool: checked,
    });
  }, [checked]);

  // useMutaionを共通化する
  const { mutate, isSuccess, isError } = useMutation(
    (body: AddStudentSchemaType) => {
      return axios.post('/students', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: {
          first_name: body.firstName,
          last_name: body.lastName,
          grade: body.grade,
          manavis_code: body.manavisCode,
          is_high_school: body.isHighSchool,
        },
      });
    },
  );

  const handleSubmit = useCallback(() => {
    const body = addStudentSchema.parse(form.values);

    mutate(body);
    if (isSuccess) {
      console.log('success');
    }
    if (isError) {
      console.log('error');
    }
  }, []);

  return (
    // TODO: zod由来のバリデーションメッセージが表示できるようにする
    <Stack align="center" spacing="lg" h={'100%'}>
      <h1>🦄New Students🎉</h1>
      <Text fz="lg">新たに生徒を登録しましょう🤩</Text>
      <Box maw={300} mx="auto" p={40}>
        <form>
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
          {/* <Select
            pb={20}
            withAsterisk
            required
            label="高校生/中学生"
            data={[
              { value: 'true', label: '高校生' },
              { value: 'false', label: '中学生' },
            ]}
            {...form.getInputProps('isHighSchool')}
          ></Select> */}

          <Checkbox
            label="高校生"
            checked={checked}
            onChange={handleChangeCheckbox}
          />

          <Group position="right" mt="md">
            <Button onClick={handleSubmit}>登録</Button>
          </Group>
        </form>
      </Box>
    </Stack>
  );
});
