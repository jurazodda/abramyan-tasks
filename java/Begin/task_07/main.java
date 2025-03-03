import java.util.Scanner;

public class main {
    public static void main(String[] args) {
        double p = 3.14;
        Scanner scanner = new Scanner(System.in);
        System.out.print("Enter the radius of the circle: ");
        double r = scanner.nextDouble();

        double l = 2 * p * r;
        System.out.printf("Circumference of the circle: %.2f\n", l);

        double s = p * r*r;
        System.out.printf("Area of the circle: %.2f\n", s);
    }
}