import java.util.Scanner;

public class main {
    public static void main(String[] args) {
        Scanner scanner = new Scanner(System.in);
        System.out.print("Enter the lenght of the rectangle: ");
        double a = scanner.nextDouble();
        System.out.print("Enter the width of the rectangle: ");
        double b = scanner.nextDouble();

        double s = a * b;
        System.out.printf("Area of the rectangle: %.2f\n", s);

        double p = 2 * (a + b);
        System.out.printf("Perimeter of the rectangle: %.2f\n", p);
    }
} 