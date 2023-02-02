function addHours(n: number): Date {
    let currentDate = new Date()
    return new Date(currentDate.getTime() + n * 60 * 60 * 1000)
}

let n = 9
let result = addHours(n)
console.log(result.toString())
