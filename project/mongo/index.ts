import { PrismaClient } from "@prisma/client"

const prisma = new PrismaClient()

async function main() {
    await prisma.test.create({
        data: {
            id: 3,
            device_id: 3,
            name: "sakamoto",
        },
    })
    const test_data = await prisma.test.findMany()
    console.dir(test_data)
}
main()
    .then(async () => {
        await prisma.$disconnect()
    })
    .catch(async (e) => {
        console.error(e)
        await prisma.$disconnect()
        process.exit(1)
    })
