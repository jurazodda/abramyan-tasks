import java.util.Scanner;

public  class main {
    public static void main(String[] args) {
        Scanner scanner = new Scanner(System.in);
        System.out.print("Enter the first number: ");
        double num1 = scanner.nextDouble();

        System.out.print("Enter the second number: ");
        double num2 = scanner.nextDouble();

        double avg = (num1 + num2) / 2;
        System.out.printf("Average value: %.2f\n", avg);
    }
}