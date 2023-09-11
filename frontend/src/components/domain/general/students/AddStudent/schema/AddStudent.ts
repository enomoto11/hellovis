import { z } from 'zod';

// manavisCode は 6桁の数字を文字列にしたもの
const manavisCodeRegex = /^[0-9]{6}$/;

// gradeは1~3の整数を文字列にしたもの
const gradeRegex = /^[1-3]$/;

export const addStudentSchema = z.object({
  manavisCode: z
    .string()
    .regex(manavisCodeRegex, { message: 'マナビスコードは6桁の数字です' })
    .min(6, { message: 'マナビスコードは6桁の数字です' }),
  firstName: z.string().nonempty({ message: '名前を入力してください' }),
  lastName: z.string().nonempty({ message: '苗字を入力してください' }),
  grade: z
    .string()
    .regex(gradeRegex, { message: '学年は1~3年生で入力してください' })
    .nonempty({ message: '学年を選択してください' }),
  isHighSchool: z.boolean(),
});

export type AddStudentSchemaType = z.infer<typeof addStudentSchema>;
