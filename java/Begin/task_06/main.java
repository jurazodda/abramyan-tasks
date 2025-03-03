
import java.util.Scanner;

public class main {
    public static void main(String[] args) {
        Scanner scanner = new Scanner(System.in);
        System.out.print("Enter the edge lenghts of the rectangular parallelepiped: ");
        double a = scanner.nextDouble();
        double b = scanner.nextDouble();
        double c = scanner.nextDouble();

        double v = a * b * c;
        System.out.printf("Volume of the rectangular parallelepiped: %.2f\n", v);

        double s = 2 * (a*b + b*c + c*a);
        System.out.printf("Surface area of the rectangular parallelepiped: %.2f\n", s);
    }
}