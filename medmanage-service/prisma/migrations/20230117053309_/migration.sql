/*
  Warnings:

  - The primary key for the `med_schedule` table will be changed. If it partially fails, the table could be left without primary key constraint.
  - You are about to drop the `user_profile` table. If the table is not empty, all the data it contains will be lost.

*/
-- AlterTable
ALTER TABLE "med_schedule" DROP CONSTRAINT "med_schedule_pkey",
ADD CONSTRAINT "med_schedule_pkey" PRIMARY KEY ("id");

-- DropTable
DROP TABLE "user_profile";

-- CreateTable
CREATE TABLE "User" (
    "id" SERIAL NOT NULL,
    "name" TEXT NOT NULL,

    CONSTRAINT "User_pkey" PRIMARY KEY ("id")
);

-- AddForeignKey
ALTER TABLE "med_schedule" ADD CONSTRAINT "med_schedule_user_id_fkey" FOREIGN KEY ("user_id") REFERENCES "User"("id") ON DELETE RESTRICT ON UPDATE CASCADE;
