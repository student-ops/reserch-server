/*
  Warnings:

  - The primary key for the `user_profile` table will be changed. If it partially fails, the table could be left without primary key constraint.
  - You are about to drop the column `id` on the `user_profile` table. All the data in the column will be lost.

*/
-- AlterTable
ALTER TABLE "user_profile" DROP CONSTRAINT "user_profile_pkey",
DROP COLUMN "id",
ADD COLUMN     "user_id" SERIAL NOT NULL,
ADD CONSTRAINT "user_profile_pkey" PRIMARY KEY ("user_id");

-- CreateTable
CREATE TABLE "med_schedule" (
    "id" SERIAL NOT NULL,
    "user_id" INTEGER NOT NULL,
    "state" BOOLEAN NOT NULL,
    "name" TEXT NOT NULL,

    CONSTRAINT "med_schedule_pkey" PRIMARY KEY ("id")
);
