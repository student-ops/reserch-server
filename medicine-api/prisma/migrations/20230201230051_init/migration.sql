-- CreateTable
CREATE TABLE "User" (
    "id" SERIAL NOT NULL,
    "name" TEXT NOT NULL,

    CONSTRAINT "User_pkey" PRIMARY KEY ("id")
);

-- CreateTable
CREATE TABLE "med_schedule" (
    "id" SERIAL NOT NULL,
    "user_id" INTEGER NOT NULL,
    "state" BOOLEAN NOT NULL,
    "taken_day" INTEGER NOT NULL,
    "taken_time" INTEGER NOT NULL,
    "tabletes" INTEGER NOT NULL,
    "name" TEXT NOT NULL,

    CONSTRAINT "med_schedule_pkey" PRIMARY KEY ("id")
);

-- AddForeignKey
ALTER TABLE "med_schedule" ADD CONSTRAINT "med_schedule_user_id_fkey" FOREIGN KEY ("user_id") REFERENCES "User"("id") ON DELETE RESTRICT ON UPDATE CASCADE;
