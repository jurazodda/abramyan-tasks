import java.util.Scanner;

public class main {
    public static void main(String[] args) {
        Scanner scanner = new Scanner(System.in);
        System.out.print("Enter the edge of the cube: ");
        double a = scanner.nextDouble();

        double v = a * a * a;
        System.out.printf("Volume of the cube: %.2f\n", v);

        double s = 6 * a * a;
        System.out.printf("Surface area of the cube: %.2f\n", s);
    }
}