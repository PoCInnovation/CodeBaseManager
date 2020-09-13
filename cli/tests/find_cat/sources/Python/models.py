from django.contrib.auth.models import User
from django.db import models
from django.urls import reverse
from django.utils import timezone
from django.conf import settings

# Create your models here.
InviteChoice = {
    (1, 'Manager'),
    (2, 'Responsable Clientèle'),
    (3, 'Responsable Facturation'),
    (4, 'Commercial'),
}

ServiceStatusChoice = {
    (0, 'Pas encore effectué'),
    (1, 'Effectué'),
    (2, 'Ne sera jamais effectué'),
}

CompanyFactuDelayChoice = {
    (15, '15 Jours'),
    (30, '30 Jours'),
    (45, '45 Jours'),
    (60, '60 Jours'),
    (75, '75 Jours'),
    (90, '90 Jours'),
}


def GetDate():
    return timezone.now().date()


def getInvoiceStorage(instance, filename):
    return 'Factures/{0}_{1}/{2}_{3}'.format(
        instance.company.name,
        instance.company.id,
        instance.id,
        instance.date
    )


class Company(models.Model):
    name = models.CharField(
        max_length=150,
        verbose_name="Nom de l'entreprise.",
        help_text="Préciser le nom de l'entreprise.",
    )
    ceo = models.OneToOneField(
        User,
        on_delete=models.CASCADE,
        verbose_name="CEO de l'entreprise.",
        help_text="Préciser le CEO de l'entreprise.",
    )
    siret = models.CharField(
        default=None,
        max_length=14,
        verbose_name="Numéro SIRET de l'entreprise.",
        help_text="Numéro SIRET de l'entreprise.",
    )
    invoice_nb = models.PositiveIntegerField(
        default=0,
        verbose_name="Numéro de factures de l'entreprise.",
        help_text="Numéro de factures de l'entreprise.",
    )
    junior_day = models.PositiveIntegerField(
        default=0,
        verbose_name="Prix du Jour homme junior de cette entreprise.",
        help_text="Précisez le prix du jour homme junior de cette entreprise.",
    )
    senior_day = models.PositiveIntegerField(
        default=0,
        verbose_name="Prix du Jour homme senior de cette entreprise.",
        help_text="Précisez le prix du jour homme senior de cette entreprise.",
    )
    facturation_delay = models.PositiveSmallIntegerField(
        default=45,
        choices=CompanyFactuDelayChoice,
        verbose_name="Délai de facturation.",
        help_text="Précisez le délai à partir duquel un client est considéré en retard de paiement.",
    )

    class Meta:
        verbose_name = "entreprise"
        verbose_name_plural = "entreprises"
        ordering = ['ceo__id']

    def __str__(self):
        return self.name


class Manager(models.Model):
    user = models.OneToOneField(
        User,
        on_delete=models.CASCADE,
        verbose_name="l'utilisateur lié à ce manager",
        help_text="préciser l'utilisateur lié à ce manager.",
    )
    company = models.ForeignKey(
        Company,
        default=None,
        on_delete=models.CASCADE,
        verbose_name="entreprise du manager.",
        help_text="préciser l'entreprise de ce manager.",
    )
    role = models.PositiveSmallIntegerField(
        default=3,
        verbose_name="rôle du manager.",
        help_text="préciser le rôle du manager.",
        choices={
            (1, 'Manager General'),
            (2, 'Account Manager'),
            (3, 'Factu Manager'),
        },
    )

    class Meta:
        verbose_name = "manager"
        verbose_name_plural = "managers"
        ordering = ['company__id', 'role', 'user', 'user__first_name', 'user__last_name']

    def __str__(self):
        return "%s %s" % (self.user.first_name, self.user.last_name)


class Commercial(models.Model):
    user = models.OneToOneField(
        User,
        default=None,
        on_delete=models.CASCADE,
        verbose_name="l'utilisateur lié à ce commercial.",
        help_text="préciser l'utilisateur lié à ce commercial.",
    )
    company = models.ForeignKey(
        Company,
        default=None,
        on_delete=models.CASCADE,
        verbose_name="l'entreprise de ce commercial.",
        help_text="préciser l'entreprise de ce commercial."
    )

    class Meta:
        verbose_name = "Commercial"
        verbose_name_plural = "Commercials"
        ordering = ['company__id', 'user', 'user__first_name', 'user__last_name']

    def __str__(self):
        return "%s %s" % (self.user.first_name, self.user.last_name)


class Client(models.Model):
    name = models.CharField(
        max_length=150,
        default=None,
        verbose_name="le nom du client.",
        help_text="le nom du client."
    )
    email = models.EmailField(
        max_length=150,
        default=None,
        verbose_name="email du client",
        help_text="l'email du client",
    )
    commercial = models.ForeignKey(
        Commercial,
        default=None,
        on_delete=models.CASCADE,
        verbose_name="le commercial ayant ramené le client.",
        help_text="préciser le commercial relatif à ce client.",
    )
    account_manager = models.ForeignKey(
        Manager,
        default=None,
        on_delete=models.CASCADE,
        verbose_name="le responsable clientièle de ce client.",
        help_text="préciserle responsable clientièle de ce client.",
    )
    company = models.ForeignKey(
        Company,
        default=None,
        on_delete=models.CASCADE,
        verbose_name="l'entreprise du client.",
        help_text="préciser l'entreprise du client.",
    )
    created_at = models.DateField(
        default=GetDate,
        editable=False,
        verbose_name="date de création du client.",
    )

    class Meta:
        verbose_name = "client"
        verbose_name_plural = "client"
        ordering = ['company__id', 'commercial__id', 'account_manager__id', 'name']

    def __str__(self):
        return self.name

    def get_absolute_url(self, comp_id):
        return reverse('mvp-client-details', args=[comp_id, str(self.id)])


class Contract(models.Model):
    description = models.CharField(
        max_length=100,
        verbose_name="description du contrat",
        help_text="description du contrat",
    )
    start_date = models.DateField(
        default=GetDate,
        verbose_name="date de début du contrat",
        help_text="date de début du contrat"
    )
    duration = models.PositiveIntegerField(
        default=1,
        verbose_name="durée totale du contrat (en mois).",
        help_text="précisez la durée totale du contrat (en mois).",
    )
    end_date = models.DateField(
        default=GetDate,
        verbose_name="date de fin du contrat",
        help_text="date de fin du contrat."
    )
    facturation = models.PositiveSmallIntegerField(
        default=1,
        verbose_name="Fréquence de la facturation.",
        help_text="Fréquence de la facturation.",
    )
    price = models.PositiveIntegerField(
        default=0,
        verbose_name="montant total du contrat.",
        help_text="montant total du contrat (EN EUROS)",
    )
    payed = models.BooleanField(
        default=False,
        verbose_name="si le contrat est payé en intégralité.",
        help_text="précisez si le contrat est payé en intégralité.",
    )
    client = models.ForeignKey(
        Client,
        default=None,
        on_delete=models.CASCADE,
        verbose_name="client cible du contrat",
        help_text="client cible du contrat",
    )
    company = models.ForeignKey(
        Company,
        default=None,
        on_delete=models.CASCADE,
        verbose_name="entreprise vendeuse du contrat.",
        help_text="entreprise vendeuse du contrat.",
    )
    commercial = models.ForeignKey(
        Commercial,
        default=None,
        on_delete=models.CASCADE,
        verbose_name="commercial à l'origine du contrat",
        help_text="commercial à l'origine du contrat",
    )
    factu_manager = models.ForeignKey(
        Manager,
        default=None,
        on_delete=models.CASCADE,
        verbose_name="le responsable de la facturation de ce contrat.",
        help_text="préciserle responsable de la facturation de ce contrat.",
        related_name="factu_manager"
    )
    validated = models.BooleanField(
        default=False,
        verbose_name="si le contract est validé.",
        help_text="précisez si le contract est validé..",
    )

    class Meta:
        verbose_name = "conseil"
        verbose_name_plural = "conseils"
        ordering = ['company__id', 'commercial__id', 'price', 'description']

    def __str__(self):
        return self.description

    def get_absolute_url(self, comp_id):
        return reverse('mvp-contract-details', args=[comp_id, str(self.id)])


class Conseil(models.Model):
    description = models.CharField(
        max_length=150,
        verbose_name="Description du conseil.",
        help_text="Description du conseil",
    )
    contract = models.ForeignKey(
        Contract,
        default=None,
        on_delete=models.CASCADE,
        verbose_name="Le contrat dans lequel est inclus ce conseil.",
        help_text="Préciser le contrat dans lequel est inclus ce conseil.",
    )
    price = models.PositiveIntegerField(
        default=0,
        verbose_name="Montant total.",
        help_text="Montant total (€)",
    )
    payed = models.BooleanField(
        default=False,
        verbose_name="Si le conseil est payé.",
        help_text="Précisez si le conseil est déjà payé.",
    )
    start_date = models.DateField(
        default=GetDate,
        verbose_name="Date de début.",
        help_text="Date de début du conseil."
    )
    duration = models.PositiveIntegerField(
        default=1,
        verbose_name="Durée totale (en mois).",
        help_text="Précisez la durée totale du conseil (en mois).",
    )
    end_date = models.DateField(
        default=GetDate,
        verbose_name="Date de fin du conseil.",
        help_text="Date de fin du conseil."
    )

    class Meta:
        verbose_name = "conseil"
        verbose_name_plural = "conseils"
        ordering = ['contract__id', 'payed', 'price', 'description']

    def __str__(self):
        return self.description

    def get_absolute_url(self, comp_id, contract_id):
        return reverse('mvp-conseil-details', args=[comp_id, contract_id, str(self.id)])


class Service(models.Model):
    description = models.CharField(
        max_length=150,
        verbose_name="Description du service",
        help_text="Description du service",
    )
    price = models.PositiveIntegerField(
        default=0,
        verbose_name="Pricing du service",
        help_text="Pricing du service (€)",
    )
    conseil = models.ForeignKey(
        Conseil,
        default=None,
        on_delete=models.CASCADE,
        verbose_name="Conseil relatif au service",
        help_text="Conseil relatif au service",
    )
    estimated_date = models.DateField(
        default=timezone.now,
        verbose_name="Date prévisionelle",
        help_text="Sélectionnez une date prévisionelle."
    )
    actual_date = models.DateField(
        default=None,
        verbose_name="Fin du service",
        help_text="Fin du service",
        null=True
    )
    payed = models.BooleanField(
        default=False,
        verbose_name="Si le service est payé.",
        help_text="Précisez si le service est déjà payé.",
    )
    junior_day = models.DecimalField(
        default=0,
        max_digits=5,
        decimal_places=2,
        verbose_name="Jour-hommes junior nécessaire pour ce service.",
        help_text="Précisez le Jour homme junior nécessaire pour ce service.",
    )
    senior_day = models.DecimalField(
        default=0,
        max_digits=5,
        decimal_places=2,
        verbose_name="Jour-hommes senior nécessaire pour ce service.",
        help_text="Précisez le Jour homme senior nécessaire pour ce service.",
    )
    done = models.SmallIntegerField(
        default=0,
        choices=ServiceStatusChoice,
        verbose_name="Si le service est effectué.",
        help_text="Précisez si le service est effectué.",
    )

    class Meta:
        verbose_name = "service"
        verbose_name_plural = "services"
        ordering = ['payed', 'estimated_date', 'price']

    def __str__(self):
        return self.description


class License(models.Model):
    description = models.CharField(
        max_length=150,
        verbose_name="Description de la licence.",
        help_text="Description de la licence."
    )
    contract = models.ForeignKey(
        Contract,
        default=None,
        on_delete=models.CASCADE,
        verbose_name="Le contrat dans lequel est inclus cette licence.",
        help_text="Préciser le contrat dans lequel est inclus cette licence.",
    )
    price = models.PositiveIntegerField(
        default=0,
        verbose_name="Montant",
        help_text="Montant (€)."
    )
    start_date = models.DateField(
        default=GetDate,
        verbose_name="Date de début.",
        help_text="Date de début."
    )
    duration = models.PositiveIntegerField(
        default=1,
        verbose_name="Durée totale (en mois).",
        help_text="Précisez la durée totale (en mois).",
    )
    end_date = models.DateField(
        default=GetDate,
        verbose_name="Date de fin.",
        help_text="Date de fin."
    )
    payed = models.BooleanField(
        default=False,
        verbose_name="Si la licence est payée.",
        help_text="Précisez si la licence est déjà payée.",
    )

    class Meta:
        verbose_name = "license"
        verbose_name_plural = "licenses"
        ordering = ['contract__id', 'price', 'description']

    def __str__(self):
        return self.description

    def get_absolute_url(self, comp_id, contract_id):
        return reverse('mvp-license-details', args=[comp_id, contract_id, str(self.id)])


class Invoice(models.Model):
    description = models.CharField(
        max_length=300,
        verbose_name="Invoice's description",
        help_text="Description de la facture",
    )
    price = models.PositiveIntegerField(
        default=0,
        verbose_name="Montant total de la facture.",
        help_text="Montant total de la facture (EN EUROS).",
    )
    date = models.DateField(
        default=GetDate,
        verbose_name="Date de la facture",
        help_text="Date de la facture."
    )
    pdf = models.FileField(
        default=None,
        verbose_name="PDF de la facture.",
        help_text="PDF de la facture.",
        null=True,
        upload_to=getInvoiceStorage,
    )
    facturated_date = models.DateField(
        default=None,
        verbose_name="Date de facturation",
        help_text="Précisez la date de facturation",
        null=True,
    )
    payed = models.BooleanField(
        default=False,
        verbose_name="Si la facture est payée.",
        help_text="Précisez si la facture est payée",
    )
    contract = models.ForeignKey(
        Contract,
        default=None,
        on_delete=models.CASCADE,
        verbose_name="Le contrat relatif à cette facture.",
        help_text="Préciser le contrat relatif à cette facture.",
    )
    company = models.ForeignKey(
        Company,
        on_delete=models.CASCADE,
        verbose_name="L'entreprise de cette facture.",
        help_text="Préciser l'entreprise de cette facture.",
    )
    conseils = models.ManyToManyField(
        Conseil,
        default=None,
        verbose_name="Le ou les conseil(s) relatif(s) à cette facture.",
        help_text="Préciser le ou les conseil(s) à facturer.",
    )
    licenses = models.ManyToManyField(
        License,
        default=None,
        verbose_name="Le ou les licenses(s) relatif(s) à cette facture.",
        help_text="Préciser le ou les licenses(s) à facturer.",
    )
    number = models.PositiveIntegerField(
        default=0,
        verbose_name="Numéro de la facture.",
        help_text="Numéro de la facture.",
    )
    payed_date = models.DateField(
        default=None,
        verbose_name="Date d'encaissement",
        help_text="Précisez la date d'encaissement",
        null=True,
    )


    class Meta:
        verbose_name = "facture"
        verbose_name_plural = "factures"
        ordering = ['company__id', 'payed', 'price', 'description']

    def __str__(self):
        return self.description

    def get_absolute_url(self, comp_id):
        return reverse('mvp-invoice-details', args=[comp_id, str(self.id)])


class Invite(models.Model):
    email = models.EmailField(
        max_length=150,
        default=None,
        verbose_name="Email du destinataire.",
        help_text="L'email du destinataire.",
    )
    company = models.ForeignKey(
        Company,
        on_delete=models.CASCADE,
        verbose_name="L'entreprise de cette invitation.",
        help_text="Préciser l'entreprise de cette invitation.",
    )
    role = models.PositiveSmallIntegerField(
        default=4,
        verbose_name="Rôle de la personne",
        help_text="Préciser le rôle de la personne",
        choices=InviteChoice,
    )

    class Meta:
        verbose_name = "invitation"
        verbose_name_plural = "invitations"
        ordering = ['company__id', 'role', ]

    def __str__(self):
        return self.email


class EmailDatabase(models.Model):
    email = models.EmailField(
        max_length=150,
        default=None,
        verbose_name="Email du destinataire.",
        help_text="L'email du destinataire.",
    )

    class Meta:
        verbose_name = "email database"
        verbose_name_plural = "emails database"
        ordering = ['email', ]

    def __str__(self):
        return self.email

