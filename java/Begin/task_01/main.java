import java.util.Scanner;

public class main {
    public static void main(String[] args) {
        Scanner scanner = new Scanner(System.in);
        System.out.print("Enter the side of the square: ");
        double a = scanner.nextDouble();

        double p = 4 * a;
        System.out.printf("Perimeter of the square: %.2f\n", p);
    }
}