/*
  Warnings:

  - The primary key for the `med_schedule` table will be changed. If it partially fails, the table could be left without primary key constraint.

*/
-- AlterTable
ALTER TABLE "med_schedule" DROP CONSTRAINT "med_schedule_pkey",
ADD CONSTRAINT "med_schedule_pkey" PRIMARY KEY ("user_id");
